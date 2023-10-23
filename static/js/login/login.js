let formData =document.querySelector(".form")
document.querySelector(".modal-close").addEventListener("click",()=>{
    modal.setAttribute('style', 'visibility:hidden;')
})
formData.onsubmit=async (e)=>{
    e.preventDefault()

    fetch('/login',{
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
                if (data.msg==="登录成功"){
                    setTimeout(function () {
                        window.location.href="/"
                    },700)
                }


            }
        )
}







