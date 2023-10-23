let list =document.querySelector(".list")


fetch('/data',)
    .then
    (response=>{return response.json()})
    .then
    (datas=>{
        console.log(datas)
        for (const data of datas){
            const Name =data.User.Name
            const postId =data.ID
            const commentNum=data.CommentNum
            const createTime=data.CreateTime
            const contentText=data.Content
            const contentTag="早起"
            const likeNum=data.LikeNum
            list.insertAdjacentHTML("afterbegin",'<li> <img src="../static/img/index/profile.PNG" class="profile" alt="头像"> <div class="content"> <div class="content-head"> <div class="content-head-id"> <div class="id-name">'+
                Name +' </div> <div class="id-major"> </div> </div> <div class="content-head-time">'+
                createTime +'</div> </div> <div class="content-body"> <div class="content-body-text"> <div class="lqduo">&ldquo;</div> <div class="textAndRdquo">'+
                contentText +'<span>&rdquo;</span></div> </div> </div> <div class="content-bottom"> <div class="content-bottom-tag"> <img src="../static/img/index/sun.svg" class="content-bottom-tag-img"> <div class="content-bottom-tag-text">' +
                contentTag +' </div> </div> <div class="content-bottom-interact" > <img src="../static/img/index/like.png" class="interact-like"> <img src="../static/img/index/liked.png" class="interact-liked"> <div class="interact-like-num">' +
                likeNum +'</div> <img src="../static/img/index/message.png" class="interact-comment"> <div>'+
                commentNum+'</div> <div class="interact-postId">'+postId+'</div></div> </div> </div> </li>'
            )


        }


        {
            const interacts = document.querySelectorAll(".content-bottom-interact")
            for (const interact of interacts) {
                const like = interact.querySelector(".interact-like")
                const liked = interact.querySelector(".interact-liked")
                const likeNum = interact.querySelector(".interact-like-num")
                const postId = interact.querySelector(".interact-postId")
                const comment =interact.querySelector(".interact-comment")
                comment.onclick=()=>{

                    window.location.href=`/comment/${postId.innerText}`
                }
                let jsonData = {
                    action: "",
                    postId: ""

                }


                jsonData.action = "query"
                jsonData.postId = postId.textContent
                liked.addEventListener("click", unlikeClick)
                like.addEventListener("click", likeClick)
                fetch('/postLike', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(jsonData)
                })
                    .then(res => res.json())
                    .then(data => {
                        if (data.isGuest) {
                            like.removeEventListener("click", likeClick)
                            liked.removeEventListener("click", unlikeClick)
                        }else if(data.isLike) {
                            liked.style.display = "block"
                            like.style.display = "none"

                        } else {
                            like.style.display = "block"
                            liked.style.display = "none"
                        }


                    })

                function likeClick() {
                    jsonData.action = "like"

                    fetch('/postLike', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify(jsonData)
                    })
                    liked.style.display = "block"
                    like.style.display = "none"
                    let likeNumValue = parseInt(likeNum.textContent)
                    likeNumValue++
                    likeNum.textContent = likeNumValue.toString()
                }

                function unlikeClick() {
                    like.style.display = "block"
                    liked.style.display = "none"
                    jsonData.action = "unlike"

                    fetch('/postLike', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify(jsonData)
                    })
                    let likeNumValue = parseInt(likeNum.textContent)
                    likeNumValue--
                    likeNum.textContent = likeNumValue.toString()

                }

            }


        }


        }
    )




    function currentTime(){
        // 获取当前日期和时间
        const currentDate = new Date();

// 获取年、月、日、小时和分钟
        const year = currentDate.getFullYear();
        const month = (currentDate.getMonth() + 1).toString().padStart(2, '0'); // 月份从0开始，需要加1
        const day = currentDate.getDate().toString().padStart(2, '0');
        const hours = currentDate.getHours().toString().padStart(2, '0');
        const minutes = currentDate.getMinutes().toString().padStart(2, '0');

// 构建日期时间字符串
        return  `${year}-${month}-${day} ${hours}:${minutes}`;

    }



