const change =document.querySelector(".head-navigation-img")
const follow =document.querySelector(".head-navigation-text-follow")
const _new =document.querySelector(".head-navigation-text-new")
follow.onclick=()=>{
    change.style.left="-19px"
    console.log(1)
}
_new.onclick=()=>{

    change.style.left="19px"
    console.log(2)
}