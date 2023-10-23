let formData =document.querySelector(".form")

formData.onsubmit=async (e)=>{
    e.preventDefault()

    fetch('/form',{
        method:"post",
        body:new FormData(formData)}
    )
    window.location.href="/"


            }


