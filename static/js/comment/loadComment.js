


let postName =document.querySelector(".post-name")
let postTime =document.querySelector(".post-time")
let postText =document.querySelector(".post-text")
let commentList =document.querySelector(".comment-list")
let commentForm=document.querySelector(".comment-form")
let commentText=document.querySelector(".comment-form-input")
let back =document.querySelector(".back")
back.onclick=()=>{
    window.location.href="/"
}
fetch(window.location.href,{method: 'POST'})
.then(res=>res.json())
.then(post=>{
    postName.innerText="楼主:"+post.User.Name
    postTime.innerText="发布时间:"+post.CreateTime
    postText.innerText="内容:"+post.Content
   for(const comment of post.Comments){
    commentList.insertAdjacentHTML("beforeend",'<li><p class=\"comment-name\">'+
        comment.User.Name+'</p><p class=\"comment-time\">'+
        comment.CreateTime+' </p><p class=\"comment-text\">'+
        comment.Content+' </p> <div class=\"comment-ID\" style=\"display: none\">'+
   comment.ID+'</div></li>')}

commentForm.onsubmit =  (e) => {
   // console.log(commentText.value)
    let jsonData = {
        postId: post.ID,
        content: commentText.value

    }
    e.preventDefault();
    fetch("/comment", {
        method: "post",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(jsonData)
    })
    window.location.reload()

   }})



