/// antlr¡4 Distance.g4 -o prDistance -package prDistance -visitor
grammar Distance; 

options {
	language = 'Go';
}


s :  list ; 

list :  
	list '(' INT ',' INT ')'	 	#L1  
	| '(' INT ',' INT ')' '(' INT ',' INT ')'#L2 ; 


INT : '-'?[0-9]+; 
WS : [ \n\r\t]+ -> skip; 