var config = {
  mode: "fixed_servers",
  rules: {
    singleProxy: {
      host: "localhost"
    }
  }
};
chrome.proxy.settings.set(
    {value: config, scope: 'regular'},
    function() {});

//var first_url = "";
//var second_url = "";
//
////var rule = {
////        conditions: [
////          new chrome.declarativeWebRequest.RequestMatcher({
////            url: { hostSuffix: 'litterfeldt.com' } })
////          ],
////        actions: [
////          new chrome.declarativeWebRequest.SetRequestHeader({name:"Host", value:"proxy.litterfeldt.com"}),
////          new chrome.declarativeWebRequest.SetRequestHeader({name:"Proxy-Host", value:"www.litterfeldt.com"}),
////          new chrome.declarativeWebRequest.RedirectRequest({redirectUrl:"proxy.litterfeldt.com"})
////        ]};
////
////chrome.declarativeWebRequest.onRequest.addRules([rule]);
//
//chrome.webRequest.onBeforeSendHeaders.addListener(
//  function(details) {
//    //console.log(getLocation(details.url).host);
//    //console.log(details.requestHeaders);
//    console.log(first_url);
//    details.requestHeaders.push({name:"Proxy-Host",value: first_url});
//    details.requestHeaders.push({name: "Host", value: "proxy.litterfeldt.com"});
//    return {requestHeaders: details.requestHeaders};
//  },
//
//  {urls: ["http://*/*", "https://*/*"]},
//
//  ["blocking", "requestHeaders"]
//);
//chrome.webRequest.onBeforeRequest.addListener(
//  function(details){
//    var url = getLocation(details.url);
//
//    if (url.host != "proxy.litterfeldt.com")
//      first_url = _.clone(url.host, true);
//    url.host = "proxy.litterfeldt.com";
//    return {redirectUrl: url.href};
//  },
//  {urls: ["http://*/*", "https://*/*"]},
//  ["blocking"]
//);
//function getLocation(href) {
//    var location = document.createElement("a");
//    location.href = href;
//    if (location.host == "") {
//      location.href = location.href;
//    }
//    return location;
//};
