# FuncyMonkey Syntax
---
```html
<program> ::== <big function> | <big statement>
<big function> ::== func <identifier> '(' {<identifier>} ')' '{' <statement> '}'
<big statement> ::== let <identifier> = <expression>;

<statement> ::== 
    if <expression> '{' <statement> '}' [else '{' <statement> '}'] |
    for <identifier> := <expression> '{' <statement> '}' |
    <big statement> | var <identifier> <otype>; | const <identifier> = <value>; |
    return <expression>;


<expression> ::== 
    <identifier> | <call function> | <prefix expression> | <infix expression> | <value>


<call function> ::== <identifier>'(' {<identifier>} ')';
<prefix expression> ::== (!|-)<expression>
<infix expression> ::== <expression>(+|-|*|/|<|>|==|!=|:=)<expression>
<value> ::== <integer> | <float> | <string> | <array> | <hash> | <boolean> | <small function>


<identifier> ::== <letter>{<letter>}
<integer> ::== <digit>{<digit>}
<float> ::== <digit>.<digit>{<digit>}
<string> ::== "<able all character>"
<array> ::== '[' {<value>} ']'
<hash> ::== '{' {<KV map>} '}'
<boolean> ::== true | false
<small function> ::== fn'(' {<identifier>} ')' '{' <statement> '}'

<able all character> ::== yes. you think it. So, that is character? all able it

<KV map> ::== <identifier> : <value>

<letter> ::= A | B | C | ... | X | Y | Z | a | b | ... | z | _
<digit> ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9

```