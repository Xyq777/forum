const logout =document.querySelector(".user-info-logout")
const userInfo=document.querySelector(".user-info-text")






/*userInfo.innerText=`当前用户:
${getCookieValue("username")}`
*/
if (!getCookieValue("username")){
    logout.innerText="登录"
    logout.addEventListener("click",()=>{
        window.location.href="/login"
    })
}else {userInfo.innerText=`当前用户：
${decodeURIComponent(getCookieValue("username"))}`
    logout.addEventListener("click",()=>{
        fetch("/logout")
        location.reload()
    })
}




function getCookieValue(cookieName) {
    let name = cookieName + "=";
    let cookieArray = document.cookie.split(';');

    for (let i = 0; i < cookieArray.length; i++) {
        let  cookie = cookieArray[i].trim();
        if (cookie.indexOf(name) === 0) {
            return cookie.substring(name.length, cookie.length);
        }
    }
    // 如果未找到指定的cookie，返回空字符串或者可以根据需要返回其他默认值
    return "";
}

// 调用函数来获取特定键的cookie值

