class Stack{
    data = [];
    
    push(data){
        this.data.push(data);
    }

    pop(){
        if(this.data.length === 0){
            return null;
        }
        const result = this.data[this.data.length - 1];
        this.data.splice(-1,1)
        return result;
    }

    peek(){
        if(this.data.length === 0){
            return null;
        }
        return this.data[0];
    }
}

export function printList() {
    const a: Stack = new Stack();
    a.push(2);
    a.push(3);
    a.pop();
    a.peek();
    console.log(a)



}
