<div align='center' style='font-size:30px'> 中山大学计算机学院本科生实验报告（2021年秋季学期）</div>

## 课程名称：区块链技术原理

## 大作业项目：基于区块链的供应链金融平台

|  姓名  |   年级专业   |   学号   |
| :----: | :----------: | :------: |
| 冼子婷 | 19级软件工程 |          |
| 廖雨轩 | 19级软件工程 |          |
| 胡文浩 | 19级软件工程 | 18346019 |

# 一、项目背景

![image-20211219193112078](D:\study\blockchain\sup\report\report\image-20211219193112078.png)

- **传统供应链金融**： 某车企（宝马）因为其造车技术特别牛，消费者口碑好，所以其在同行业中占据绝对优势地位。因此，在金融机构（银行）对该车企的信用评级将很高，认为他有很大的风险承担的能力。在某次交易中，该车企从轮胎公司购买了一批轮胎，但由于资金暂时短缺向轮胎公司签订了 1000 万的应收账款单据，承诺 1 年后归还轮胎公司 1000 万。这个过程可以拉上金融机构例如银行来对这笔交易作见证，确认这笔交易的真实性。在接下里的几个月里，轮胎公司因为资金短缺需要融资，这个时候它可以凭借跟某车企签订的应收账款单据向金融结构借款，金融机构认可该车企（核心企业）的还款能力，因此愿意借款给轮胎公司。但是，这样的信任关系并不会往下游传递。在某个交易中，轮胎公司从轮毂公司购买了一批轮毂，但 由于租金暂时短缺向轮胎公司签订了 500 万的应收账款单据，承诺 1 年后归还轮胎公司 500 万。当轮毂公司想利用这个应收账款单据向金融机构借款融资的时候，金融机构因为不认可轮胎公司的还款能力，需要对轮胎公司进行详细的信用分析以评估其还款能力同时验证应收账款单据的真实性，才能决定是否借款给轮毂公司。这个过程将增加很多经济成本，而这个问题主要是由于该车企的信用无法在整个供应链中传递以及交易信息不透明化所导致的。 
- **区块链+供应链金融**： 将供应链上的每一笔交易和应收账款单据上链，同时引入第三方可信机构来确认这些信息的交易，例如银行，物流公司等，确保交易和单据的真实性。同时，支持应收账款的转让， 融资，清算等，让核心企业的信用可以传递到供应链的下游企业，减小中小企业的融资难度。

<div style="page-break-after:always;"></div>

# 二、方案设计

# 项目要求

- 功能一：实现采购商品—签发应收账款交易上链。例如车企从轮胎公司购买一批轮胎并签订应收账款单据。 

- 功能二：实现应收账款的转让上链，轮胎公司从轮毂公司购买一笔轮毂，便将于车企的应 收账款单据部分转让给轮毂公司。轮毂公司可以利用这个新的单据去融资或者要求车企到期时归还钱款。 

- 功能三：利用应收账款向银行融资上链，供应链上所有可以利用应收账款单据向银行申请融资。 

- 功能四：应收账款支付结算上链，应收账款单据到期时核心企业向下游企业支付相应的欠款。

# 具体设计

## 系统流通资源

- 信用值（credit）：由于信用本身不可衡量，因此在系统中使用**信用值**可以度量的概念来表示企业的可信用程度。简单来说，信用值表示这个企业在规定有限时间内能够偿还的估计金额。信用值的使用需要经过收款方的同意，才能够生成一笔与信用值相关的账单。
- 资金（funding）：资金是一个企业或者银行能够在这个系统中自由使用的资源，其可以由另外的系统，将一定的法定货币如人民币、美金等转换为其中的等价的资金值。

## 系统角色

- 系统管理员（administrator）：对整个系统进行管理，控制整体信用值的分发以及监控企业之间的借贷情况。在实际情况中，系统管理员一般由政府等监管机构担任。考虑到这类监管机构拥有其他强制手段来进行系统中的交易控制，因此在系统中只给管理员提供一些更高级的查询功能，但是将不会参与或者干涉其中的交易情况。

  ```solidity
  struct Administrator {
  	address addr;			// 管理员地址
  	uint256 creditProvided;	// 管理员发放的信用值数目
  }
  ```

- 银行（bank）：为了简化概念和方便实现，在我们的系统中的金融机构仅使用银行作为代表。作为金融机构，银行可以对企业进行认证和信用值分发的操作

  ```solidity
  struct Bank {
  	address addr;		// 地址
  	string 	name;		// 名称
  	uint256 credit;		// 可发放信用值
  	uint256 funding;	// 可用资金
  }
  ```

- 企业（company）：企业可分为以下两种类型
  - 核心企业（core company）：核心企业在金融机构中有相当的信用程度，金融机构可认为其有能力进行还贷的能力，因此可以允许其发起的融资请求
  - 普通企业（normal company）：普通企业的信用程度仍未达到金融机构的认可标准，暂时不可向金融机构发起融资请求。若后续该企业得到金融机构的认可，也可以升级为核心企业。
  
  ```solidity
  struct Company {
  	address addr;			// 地址
  	string 	name;			// 名称
  	uint256 companyType;	// 企业类型
  	uint256 credit;			// 信用值
      uint256 funding;		// 可用资金
  }
  
  uint256 companyTypeNormal = 0;	// 普通企业
  uint256 companyTypeCore = 1;	// 核心企业
  ```

## 存储设计

### Table.sol

FISCO-BCSO 中提供了 `Table.sol` 的合约，这个合约可以作为分布式数据库使用，它实现并提供了数据库的基本 CRUD 功能。因此直接在我们系统中直接引用这个合约，直接作为我们的存储系统。接下来只需要设计相对应的存储结构即可，无需额外部署数据库对地址信息、信用值等账户信息或者交易、账单等进行记录。

为了实现系统功能，使用其中的 TableFacotry 创建以下记录表

```solidity
TableFactory tf = TableFactory(0x1001);
tf.createTable(bankTable, "1", "addr,addrStr,name,credit,funding");
tf.createTable(companyTable, "1", "addr,addrStr,name,companyType,credit,funding");
tf.createTable(txTable, "1", "txID,from,fromStr,to,toStr,amount,message,txType,txState,billID");
tf.createTable(billTable, "1", "billID,from,fromStr,to,toStr,amount,createdDate,endDate,message,lock,billState,billType");

// 当做索引
tf.createTable(roleTable, "addr", "role");
```

## 交易事件上链

合约中定义了以下 `event`， `event` 为区块链对合约操作进行上链的接口，使用 `emit` 进行调用即可让区块链节点对这个事件打包进入到区块中并广播和共识，这个事件记录到链上。  

```solidity
// 需要写入到区块链的事件
event Registration(address operatorAddr, address addr, string role);
event ProvideCredit(address operatorAddr, address addr, uint256 amount);
event ProvideFunding(address operatorAddr, address addr, uint256 amount);
event WithdrawCredit(address operatorAddr, address addr, uint256 amount);
event Financing(address operatorAddr, address bankAddr, bool useBill, string message);
event ConfirmFinancing(address bankAddr, uint256 txID, bool accepted);
event Repay(address operatorAddr, address addr, uint256 amount);
event TransferBill(address operatorAddr, address from, address newTo, uint256 billID);
event TransferFunding(address operatorAddr, address to, uint256 amount);	
```

## 交易与账单结构

### 交易

交易在系统中会对关键操作进行记录，可用于提供接口给监管机构进行操作查询。

```solidity
struct Transaction {
    uint256 txID;       	// 交易 ID
    address from;	    	// 付款人地址
    address to;		    	// 收款人地址
    uint256 amount;     	// 交易总额
	string	message;		// 对该交易的一些额外备注信息
	
    uint256 txType;     	// 交易类型
	uint256 txState;		// 交易状态
	int256 billID;			// 使用账单融资时所关联的 billID
}

uint256 txTypeNormal = 0; 			// 正常交易
uint256 txTypeCreditFinacing = 1; 	// 使用信用点融资
uint256 txTypeBillFinacing = 2; 	// 使用账单融资


uint256 txStatePending = 0;		// 正在处理
uint256 txStateRefused = 1;		// 拒绝本次交易
uint256 txStateAccepted = 2;	// 接收本次交易
```

### 账单

账单仅限于被接受的交易，相当于交易双方签订成功的交易，未还账单才表示双方之间存在借账关系。

```solidity
struct Bill {
	uint256 billID;			// 账单 ID
    address from;	    	// 付款人地址
    address to;		    	// 收款人地址
    uint256 amount;			// 账单额
    string 	createdDate;	// 创建日期
    string 	endDate;		// 还款日期
    string	message;		// 账单备注信息
    
    uint256 lock;			// 是否锁定，表示该账单是否用于进行转移操作
    uint256 state;			// 账单状态
    uint256 billType;		// 账单类型，只有付款人为核心企业才能用于融资
}

uint256 billStateUnpaid = 0;	// 账单未还
uint256 billStatePaid = 1;		// 账单已还

uint256 billUnlocked = 0;
uint256 billLocked = 1;

uint256 billTypeNormal = 0; // 普通账单
uint256 billTypeCore =1;	// 核心企业为付款人的账单
```

# 数据流图

# 核心代码

# 界面展示