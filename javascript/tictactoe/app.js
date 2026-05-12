let boxes=document.querySelectorAll(".box");
let resetBtn=document.querySelector("#reset-btn");
let turnO= true;
const winPatterns=[[0,1,2],
[0,3,6],
[0,4,8],
[1,4,7],
[2,5,8],
[2,4,6],
[3,4,5],
[6,7,8],
];
const disableAllBoxes=()=>{
    for(let box of boxes){
        box.disabled=true;
    }

};
const enableBoxes=()=>{
    for(let box of boxes){
        box.disabled=false;
        box.innerText="";
    }
    turnO=true;
};

boxes.forEach((box) => {
    box.addEventListener("click",()=>{
        if(turnO){
           box.innerText="O";
           turnO=false;
        }else{
            box.innerText="X";
           turnO=true;
        }
        box.disabled=true;
        checkWinner();
    });
});

function checkWinner(){
    for(let pattern of winPatterns){
        let[a,b,c]=pattern;
        let val1=boxes[a].innerText;
        let val2=boxes[b].innerText;
        let val3=boxes[c].innerText;

        if(val1!==""&&val2!==""&&val3!==""){
            if(val1===val2&&val2===val3){
            console.log("Winner!",val1);
            alert(`${val1} wins!`);
            disableAllBoxes();
            return;
        }
    }
}
}
resetBtn.addEventListener("click",()=>{
        enableBoxes();
});

