let modalText=document.querySelector(".modal-text")
let modal =document.querySelector(".modal")
let img=document.querySelector("img")
let username =document.getElementById("Username")
username.onfocus=()=>{
    img.setAttribute('style', 'transform: translate(-60px);')
}
username.onblur=()=>{
    img.setAttribute('style', 'transform: translate(60px);')

}
