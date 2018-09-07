var ws3 = new WebSocket("ws://localhost:3000/v3/ws");
ws3.addEventListener("message", function(e){
    console.log("I am got", e)
})