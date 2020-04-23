package treesAndGraphs;

public class BSTNode extends Node {
	int aux = 0;

	BSTNode(int value) {
		super(value);
	}

	public int getAux() {
		return aux;
	}

	public void setAux(int aux) {
		this.aux = aux;
	}
}
