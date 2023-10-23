let text=document.querySelector(".text")
let tool=document.querySelector(".tooltip div")


let formData =document.querySelector(".form")
let password =document.getElementById("Password")
password.onfocus=function (){


    tool.setAttribute('style', 'visibility: visible;')
}
password.onblur=function(){
    tool.setAttribute('style', 'visibility: hidden;')
}


document.querySelector(".modal-close").addEventListener("click",()=>{
    modal.setAttribute('style', 'visibility:hidden;')
})

formData.onsubmit=async (e)=>{
    e.preventDefault()

    fetch('/reg',{
        method:"post",
        body:new FormData(formData)}
    )
        .then
        (response=>{
            return response.json()}
        )
        .then
        (data=>{

                modalText.innerText=data.msg
                modal.setAttribute('style', 'visibility: visible;')


            }
        )
}






