// let btn1 = document.querySelector("#btn1");

// // btn1.onclick=(evt)=>{
// //     console.log(evt);
// //     console.log(evt.type);
// //     console.log(evt.target);
// //     console.log(evt.clientX,evt.clientY);

// // };

// btn1.addEventListener("click",(evt)=>{
//     console.log("button1 was clicked");
//     // console.log(evt);
//     // console.log(evt.type);
// });

// btn1.addEventListener("click",()=>{
//     console.log("button1 was clicked -handler2");
// });

// const handler3=()=>{
//     console.log("button1 was clicked -handler3");
// };
// btn1.addEventListener("click",handler3);

// btn1.addEventListener("click",()=>{
//     console.log("button1 was clicked -handler4");
// });

// btn1.removeEventListener("click",handler3);

// // let div=document.querySelector("div");
// // div.onmouseover=()=>{
// //     console.log("you are inside the div");
// // };

// let modeBtn=document.querySelector("#mode");
// let currMode="light";//dark
// modeBtn.addEventListener("click",()=>{
//     if (currMode==="light"){
//         currMode="dark";
//         document.querySelector("body").classList.remove("light");
//         document.querySelector("body").classList.add("dark");
//     }else{
//         currMode="light";
//         document.querySelector("body").classList.remove("dark");
//         document.querySelector("body").classList.add("light");
//     }
//     console.log(currMode);
// });

// const employee={
//     calcTax(){
//         console.log("tax rate is 10 %");
//     }
// };

// const karan={
//     salary:5000,
// };

// karan.__proto__=employee;

// class Car{
//     start(){
//         console.log("start");
//     }
//     stop(){
//         console.log("stop");
//     }
//     setBrand(brand){
//         this.brand=brand;
//     }
// }

// let fortuner=new Car();
// fortuner.setBrand("fortuner");


// class Person{
//     constructor(name){
//         this.species="homo sapiens";
//         this.name=name;
//     }
//     seat(){
//         console.log("eat");
//     }
// }

// class Engineer extends Person{
//     constructor(name){
//         super(name);
//     }
//     work(){
//         super.seat();
//         console.log("solve problem, build something");
//     }
// }
// let engObj=new Engineer("kuldeep");
// console.log(engObj)
// let DATA="secret information";

// class User{
//     constructor(name,email){
//         this.name=name;
//         this.email=email;
//     }

//     viewData(){
//         console.log("data= ", DATA);
//     }
// }

// class Admin extends User{
//     constructor(name,email){
//         super(name,email);
//     }
//     editData(){
//         DATA="some new value";
//     }
// }

// let stud1=new User("jill","jill@example.com");
// let stud2=new User("kack","jack@example.com");
// console.log(stud1,stud2)

// let admin= new User("admin","admin@example.com");
// console.log(admin)



const URL="https://api.thecatapi.com/v1/images/search?limit=10&breed_ids=beng&api_key=REPLACE_ME"
// document.addEventListener("DOMContentLoaded", () => {
    const factPara=document.querySelector('#fact');
    const btn=document.querySelector('#btn');

    // const getFact=async()=>{
    //     let response=await fetch(URL);
    //     let data=await response.json();
    //     factPara.innerText= data.fact;
    // }

    function getFact(){
        fetch(URL).then((respanse)=>{
            return respanse.json();
        })
        .then((data)=>{
            console.log(data);
        });
    }

    btn.addEventListener("click", getFact);
// });