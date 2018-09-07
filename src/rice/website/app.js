document.getElementById("text").innerHTML = "Hello World<br /><p>Created by Nic Raboy</p>";
var ws3 = new WebSocket("ws://localhost:3000/v3/ws");
// ws3.addEventListener("message", function(e){
//     console.log("I am got", e)
// })

ws3.onmessage = function(e) {
    console.log("got data:" + e.data);
};
ws3.onerror = function(e) {
    console.log("got error:", e);
};
ws3.onclose = function(e) {
    console.log("got close:", e);
};