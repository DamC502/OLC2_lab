grammar HexExp;

options {
    language = 'Go';
}

s : ( l ','?)+ //visitor antes de entrar
    ;

l : a '.' a #L1
    | a  #L2
    ;

a : (CHAR)+ ;


CHAR : [0-9] | [a-zA-Z];