
const modal =document.querySelector(".modal")
for (const imgDiv of imgDivs ){
        const imgs =imgDiv.getElementsByTagName("img")
        for (let img of imgs){



             img.addEventListener("click",()=>{
                let imgClick=new Promise(function (resolve ,reject){
                    modal.style.display="block"
                   document.body.style.overflow="hidden"
                 const newImg=img.cloneNode(false)
                 newImg.className="insert"

                 document.body.appendChild(newImg)


                 resolve(newImg)})
                 imgClick.then((newImg)=>{
                     newImg.addEventListener("click",()=>{
                         modal.style.display="none"
                         document.body.style.overflow="auto"
                         newImg.remove()

                     })

                     console.log(1)
                 }



                )







        })



}}