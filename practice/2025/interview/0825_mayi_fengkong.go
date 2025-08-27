package interview


// 题目一
// 给出一个m*n的矩阵，要求按照螺旋顺序打印出矩阵中的每一种元素，其中1 <= m <= 10, 1 <= n <= 10
// 对于下面例子中的矩阵，打印结果为123698745
// 123
// 456
// 789

func spiralOrder(a [][]int) []int {
  m, n := len(a), len(a[0])
  u, d := 0, m-1
  l, r := 0, n-1

  ans := make([]int, 0, m*n)
  for {
    if len(ans) == m*n {
      break
    }
    for j := l; j <= r; j++ {
      ans = append(ans, a[u][j])
    }
    for i := u+1; i <= d; i++ {
      ans = append(ans, a[i][r])
    }
    if d-1 > u {
      for j := r-1; j >= l; j-- {
        ans = append(ans, a[d][j])
      }
    }
    if r-1 > l {
      for i := d-1; i > u; i-- {
        ans = append(ans, a[i][l])
      }
    }
    u++
    r--
    r--
    l++
  }
  return ans
}

func main() {
  cases := [][][]int{
    {{1,2,3},{4,5,6}},
    {{1,2}},
    {{1,2,3,4,5},{6,7,8,9,10},{11,12,13,14,15},{16,17,18,19,20}},
    {{1,2,3,4,5},{6,7,8,9,10},{11,12,13,14,15},{16,17,18,19,20},{21,22,23,24,25}},
     {{1,2,3,4,5},{6,7,8,9,10},{11,12,13,14,15},{16,17,18,19,20},{21,22,23,24,25},{26,27,28,29,30}},
  }
  expects := [][]int{
    {1,2,3,6,5,4},
    {1,2},
    {1,2,3,4,5,10,15,20,19,18,17,16,11,6,7,8,9,14,13,12},
    {1,2,3,4,5,10,15,20,25,24,23,22,21,16,11,6,7,8,9,14,19,18,17,12,13},
    {1,2,3,4,5,10,15,20,25,30,29,28,27,26,21,16,11,6,7,8,9,14,19,24,23,22,17,12,13},
  }
}

// 题目二
// 要求用有限状态机实现简单的正则表达式匹配，正则表达式支持'.' 和'*'两种语法，'.'表示任意字符，'*'表示任意次匹配，要求完全匹配，p为pattern，s为待匹配字符串，s由小写字母和数字组成
//  正则表达式测试可以使用：https://regex101.com/
// state可以通过(p_index, s_index)来表达，表示当前该匹配进行到了p的第p_index位，s的第s_index位
// 例子：
// p=abc* s=abcccc 完全匹配
// p=abc* s=abccccd 不完全匹配
// p=abc*. s=abc 完全匹配
// p=.* s=abcdd 完全匹配
// p=.*d s=abcccc 不完全匹配
// p=.*d s=abccccd 完全匹配
// class Solution:
//   def isMatch(self, s: str, p: str) -> bool: 
  
// 题目三
// 实现一个函数,接收一个字符串作为参数,该字符串是一个合法的算术表达式,可能包含加减乘除和括号。函数需要返回该表达式的计算结果。
// 例如:
// 输入: "1 + 2 * 3"
// 输出: 7
// 输入: "(1 + 2) * 3"
// 输出: 9
// 输入: "1 + (2 * 3)"
// 输出: 7

// 简单版实现：
// 用栈方式实现

// 通用版实现：
// 1. 将输入的表达式字符串通过lexer转换成一个token流，进行基础的词法分析
// 2. 用一个parser(语法分析器)来解析这个token流，最终生成一个抽象语法树(AST)
// 3. 最后解释器遍历这个AST，进行求值即可
// 复杂版不要求写出完整的代码，需要尽可能实现体现上述功能的骨架代码，包括
// ● token流枚举类型，词法分析的过程伪代码
// ● ast结构，语法分析的过程伪代码，重点体现出运算符优先级
// ● 解释执行ast的伪代码
// 更倾向于用通用版实现

// lexery.go
type TokenType int
const (
  TTDigit = 0
  TTOpPlus = 1
  TTOpMinus = 2
  TTOpMulti = 3
  TTOpDiv = 4
  TTOpLP = 5
  TTOpRP = 6
  TTSpace = 7
)

type Token struct {
  typ TokenType
  //rawVal []byte
  intVal int
}

type lexer struct {
  src []byte
  idx int
}

func NewLexer(s []byte) &lexer {
  return &lexer{src: s}
}

func (l *lexer) GetTokens() ([]*Token, error) {
  return l.GetExpr()
}

func (l *lexer) GetExpr() ([]*Token, error) {
  l.ScanWhile(TTSpace)
  switch l.next() {
    case TTOpLP:
      if err := l.GetExpr(); err != nil {
        // error
      }
      if err := l.GetOp(TTOpRP); err != nil {
        // error
      }
    case TTDigit:
      if l.hasMore() {
        if err := l.GetSomeOp([]TTType{TTOpPlus, TTOPMius, TTOpMulti, TTOpDiv}); err != nil {
          // error
        }
        if err := l.GetDigit(); err != nil {
          // error
        }
      }
    default:
    // invalid, error
  }
}

func (l *lexer) GetDigit() *Token {
}

func (l *lexer) GetOp() *Token {
}

func (l *lexer) ScanWhile(t TokenType) {}

func (l *lexer) next() {}

// parser.go
type ASTNodeType int
const (
  ASTNodeTypeLiteral = 0
  ASTNodeTypeBinary = 1
  ASTNodeTypeExpr = 2
)
type ASTNode struct {
  typ ASTNodeType
  intVal int
  children []*ASTNode
}

type parser struct {
  tokens []*Token
  nested int
  rpStack []int
  preType
}





func  Parse(tokens []*Token) (Expression, error) {
  parser := &parser{}
  preParse()
  expr, err := p.GetExpr(tokens)
  if err != nil {

  }
  return 
}

func (p *parser) GetExpr(tokens []*Token) (Expression, error) {
  if len(tokens) == 0 {
    return &Expression{} // empty
  }
  next := 0
  if tokens[next].typ == TokenTypeOpLP {
    p.nested++
    idx, err := p.GetRp(p.nested)
    if err != nil {
      // error case
    }
    next = idx
    expr := &Expression{
      typ: ASTNodeTypeExpr,
    }
    innerExpr, err := p.GetExpr(p.tokens[next+1:idx])
    if err != nil {}
    expr.children = innerExpr
    
  } else if p.tokens[next].typ == TokenTypeDigit {
    lExpr := ToLiteralExpression(p.tokens[0])
    op, err := p.MustPeek(TokenTypeOp)
    if err != nil {

    }
    if preType >= op {
      expr := &Expression{
        left: preExpr,
        right: lExpr,
      }
    } else {
      rExpr := p.GetExpr()
      expr := &Expression{
        left: lexpr,
        op: op,
        right: rExpr,
      }
    }
    
  } else {
    // error
  }
  if 
}

3 * (5 * 2 + (4 - 1 + func1()))

Expression {
  eval() 
}

Number extends Expression {
  eval {
    return 
  }
}

{
  left: Expression
  op: xxx
  right: Expression
}

