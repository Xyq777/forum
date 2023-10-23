const backTop = document.querySelector(".back-top");

window.addEventListener("scroll", () => {

    if (window.scrollY > 200) {
        backTop.style.bottom = "20px"

    } else {

        backTop.style.bottom = "-30px"

    }
});backTop.addEventListener("click",()=>{
    window.scrollTo(
        {
            top:0,
            behavior:"smooth"
        }
    )


})