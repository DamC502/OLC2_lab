package analizador;

public class NoTerminal {
    Double decimal = 0.0;
    int entero = 0;
    int nivel = 0;

    public NoTerminal(Double decimal, int entero, int nivel) {
        this.decimal = decimal;
        this.entero = entero;
        this.nivel = nivel;
    }

    public int getNivel() {
        return nivel;
    }

    public void setNivel(int nivel) {
        this.nivel = nivel;
    }

    public void setDecimal(Double decimal) {
        this.decimal = decimal;
    }

    public void setEntero(int entero) {
        this.entero = entero;
    }

    public Double getDecimal() {
        return decimal;
    }

    public int getEntero() {
        return entero;
    }
}
