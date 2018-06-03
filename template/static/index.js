document.addEventListener("DOMContentLoaded", function (event) {

    updateProxy();
    proxyButton();

    setInterval(function () {
        proxyButton();
    }, 500);

    setInterval(function () {
        updateProxy();
    }, 30000);

});

function updateProxy() {
    fetch('/json', {
        method: 'POST',
        body: new URLSearchParams("password=test")
    })
        .then(res => res.json())
        .then((out) => {
            document.getElementById('num').innerHTML = out.Proxies.length + ' proxies';
            document.getElementsByTagName('textarea')[0].innerHTML = out.Proxies.join("\n");
        })
        .catch(err => { throw err; });
}

function proxyButton() {
    fetch('/json', {
        method: 'POST',
        body: new URLSearchParams("password=test")
    })
        .then(res => res.json())
        .then((out) => {
            randomNumber = Math.floor((Math.random() * out.Proxies.length) + 0);
            randomProxy = out.Proxies[randomNumber].split(":");
            document.getElementById('b').setAttribute("href", 'tg://socks?server=' + randomProxy[0] + '&port=' + randomProxy[1]);
        })
        .catch(err => { throw err; });
}