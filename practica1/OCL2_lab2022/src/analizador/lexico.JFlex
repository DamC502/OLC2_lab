package analizador;
import java_cup.runtime.Symbol;

%%
%class Lexico
%public
%line
%char
%cup
%unicode
%ignorecase

%init{
    yyline = 1;
    yychar = 1;
%init}


BLANCOS=[ \r\t]+
l= [a-zA-AñÑ]+
%%


{BLANCOS} {}
{l} {return new Symbol(sym.TK_PALABRA, yyline, (int) yychar, yytext());}
"1" {return new Symbol(sym.TK_UNO, yyline, (int) yychar, yytext());}
"0" {return new Symbol(sym.TK_CERO, yyline, (int) yychar, yytext());}
"," {return new Symbol(sym.TK_COMMA, yyline, (int) yychar, yytext());}


. {
    System.out.println("Este es un error lexico: "+yytext()+
    ", en la linea: "+yyline+", en la columna: "+yychar);
}