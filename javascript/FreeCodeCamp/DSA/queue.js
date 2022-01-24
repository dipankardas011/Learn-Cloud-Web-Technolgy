
const que = [1, 2, 3, 24, 6, 7, 7, 8];

__main__();

class Queue { 
    constructor() {
        let que = [];
    }

    __push(num) {
        que.push(num);
    }

    __peek() {
        return que[0];
    }

    __pop() {
        que.pop();
    }
};

function __main__() {
    // console.log("Queue ", que);
    // __push(1232);
    // console.log("Queue ",que);
    // console.log("Popped element: ", __pop());

    
}