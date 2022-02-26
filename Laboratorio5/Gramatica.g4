grammar Gramatica;
/*
E		E + E
	|	E – E
	|	E * E
	|	E / E
	|	( E )
	|	NUM

*/

options {
    language='Go';
}

@parser::header {
    import  _"Laboratorio5/generador"
    import "Laboratorio5/entorno"
}

@parser::members{
    type Nodo struct {
        dir string
        tipo string
        num int
    }

    var desp int = 0
    var temp int = 1

    // var env []*entorno.Entorno
    var sup *entorno.Entorno

    // func push(e *entorno.Entorno) {
    //     env = append(env, e)
    // }

    // func pop() *entorno.Entorno {
    //     if len(env) < 1 {
    //         panic("empty env")
    //     }
    //     result := env[len(env) - 1]
    //     env = env[:len(env) - 1]
    //     return result
    // }

    func gen(out string) {
        fmt.Println(out);
    }
}

start
    : marcador dec  { entorno.Mostrar(sup) }
    | exp EOF
    ;

marcador
    : { sup = entorno.NewEntorno(nil); desp = 0; }
    ;

dec
    : t=tipo id=ID                  {
                                        sim := entorno.NewSimbolo($id.text, $t.tipoDato, desp)
                                        sup.Put($id.text, sim)
                                        desp = desp + 1
                                    }
        ';' dec
    | 
    ;

tipo returns[ string tipoDato ]
    : t='int'                       {
                                        $tipoDato = $t.text
                                    }
    ;

exp returns [ *Nodo nodo ]
    : e1=exp op=('*'|'/')  e2=exp    {
                                        $nodo = &Nodo { }
                                        temp = temp - $e1.nodo.num - $e2.nodo.num
                                        //  $nodo.dir = generador.NewTemp()
                                        $nodo.dir = "t" + strconv.Itoa(temp)
                                        cad := $nodo.dir + " = " + $e1.nodo.dir + " " + $op.text + " "
                                        cad += $e2.nodo.dir
                                        gen(cad)
                                        $nodo.num = 1
                                        temp = temp + 1
                                    }
    | e1=exp  op=('+'|'-') e2=exp    {
                                        $nodo = &Nodo { }
                                        temp = temp - $e1.nodo.num - $e2.nodo.num
                                        $nodo.dir = "t" + strconv.Itoa(temp)
                                      //  $nodo.dir = generador.NewTemp()
                                        cad := $nodo.dir + " = " + $e1.nodo.dir + " " + $op.text + " "
                                        cad += $e2.nodo.dir
                                        gen(cad)
                                        $nodo.num = 1
                                        temp = temp + 1
                                    }
    | '(' e=exp ')'                 {
                                        $nodo = &Nodo { }
                                        $nodo.dir = $e.nodo.dir
                                        $nodo.num = $e.nodo.num
                                    }
    | n=NUM                         {
                                        $nodo = &Nodo { }
                                        $nodo.num = 0
                                        $nodo.dir = $n.text
                                    }
    ;


// Tokens
MAS: '+';
MEN: '-';
POR: '*';
DIV: '/';
PARI: '(';
PARD: ')';
NUM: [0-9]+;
ID: [_A-Za-z][_A-Za-z0-9]*;
WHITESPACE: [ \r\n\t]+ -> skip;