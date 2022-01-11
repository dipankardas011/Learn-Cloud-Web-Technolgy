
const que = [1, 2, 3, 24, 6, 7, 7, 8];

__main__();

function __push(a) {
    que.push(a);
}

function __pop() {
    return que.shift();
}


function __main__() {
    console.log("Queue ", que);
    __push(1232);
    console.log("Queue ",que);
    console.log("Popped element: ", __pop());
}