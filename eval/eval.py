class Evaluate(object):
    def __init__(self, formula):
        self.ops = [] # 算子栈
        self.vals = [] # 算符栈

        if formula.strip() != "":
            for s in formula: # 在不支持字符串字符迭代的语言中, 应用队列实现
                if s == "(":
                    pass
                elif s == "+":
                    self.ops.append(s)
                elif s == "-":
                    self.ops.append(s)
                elif s == "*":
                    self.ops.append(s)
                elif s == "/":
                    self.ops.append(s)
                elif s == ")": # 计算括号内表达式
                    op = self.ops.pop()
                    v = self.vals.pop()
                    if op == "+":
                        v += self.vals.pop()
                    elif op == "-":
                        v = self.vals.pop() - v
                    elif op == "*":
                        v *= self.vals.pop()
                    elif op == "/":
                        v = self.vals.pop() / v
                    self.vals.append(v)
                else: # 非算符, 即算子, 则压入算子栈
                    self.vals.append(float(s))

        print(self.vals.pop())


evaluate = Evaluate("((3*4)+1)")