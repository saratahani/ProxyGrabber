$(function () {

    updateProxy();
    proxyButton();

    setInterval(function() {
        updateProxy();
    }, 60 * 1000);

    setInterval(function() {
        proxyButton();
    }, 1 * 1000);

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