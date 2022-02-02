package com.lf;


import analizador.Lexico;
import analizador.Parser;

import java.io.FileNotFoundException;
import java.io.FileReader;

public class Main {

    public static void main(String[] args) {
	// write your code here

        try {
            Parser syntaxParser = new Parser(new Lexico(new FileReader("entrada.txt")));
            syntaxParser.parse();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (Exception e) {
            e.printStackTrace();
        }

    }
}
