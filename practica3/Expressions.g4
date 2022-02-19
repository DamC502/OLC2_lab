/// antlr4 Expressions.g4 -o prExpr -package prExpr -visitor
grammar Expressions; 

options {
	language = 'Go';
}

s : (expr NL?)+ ; 
	

expr : expr '*' expr 	#Mul
	| expr '+' expr 	#Add
	| '(' expr  ')'		#Parens
	| INT				#Int
	| DECIMAL			#Decimal
	; 

DECIMAL: [0-9]+('.'|',')[0-9]+; 
INT: [0-9]+ ; 
WS: [ \t]+ -> skip ; 
NL: '\r'?'\n';
