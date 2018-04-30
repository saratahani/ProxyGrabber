$(function () {

    updateProxy();
    proxyButton();

    setInterval(function() {
        proxyButton();
    }, 500);

    setInterval(function() {
        updateProxy();
    }, 30000);

});

function updateProxy(){
    $.getJSON("/json", function (data) {
        $("#num").html(data.Proxies.length + " proxies");
        $("textarea").html(data.Proxies.join("\n"));
    }
);}

function proxyButton(){
    $.getJSON("/json", function (data) {
        randomNumber = Math.floor((Math.random() * data.Proxies.length) + 0);
        randomProxy = data.Proxies[randomNumber].split(":");
        $("#b").attr("href", 'tg://socks?server=' + randomProxy[0] + '&port=' + randomProxy[1]);
    }
);}