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
%%


{BLANCOS} {}
"1" {return new Symbol(sym.TK_UNO, yyline, (int) yychar, yytext());}
"0" {return new Symbol(sym.TK_CERO, yyline, (int) yychar, yytext());}
"." {return new Symbol(sym.TK_DOT, yyline, (int) yychar, yytext());}


. {
    System.out.println("Este es un error lexico: "+yytext()+
    ", en la linea: "+yyline+", en la columna: "+yychar);
}