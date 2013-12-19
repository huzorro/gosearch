/**
 * Created by huzorro on 13-12-12.
 */

$(document).ready(function () {
    $("#bt_search").click(function (e) {
        e.preventDefault();
        $("#results").empty();
        var query = $("#query").val();
        if (query) {
            search(query);
        }
    });

    function search(query) {
        var qd = {Query: query};
        $.getJSON("/search", qd, function (data) {
            var items = data.d.results;
            for (var k = 0, len = items.length; k < len; k++) {
                var item = items[k];
                show(item, query);
            }
        });
    }

    function show(item, query) {
        var li = document.createElement("li");

        var h3  = document.createElement("h3");
        var a = document.createElement("a");
        a.href = item.Url;
        a.target = "_blank";

        var reg = new RegExp(query, "ig");
        var dest = "<span class='highlight'>".concat(query, "</span>");

        $(a).append(item.Title.replace(reg, dest));

        var desc = document.createElement("p");
        desc.className = "outline";
        $(desc).append(item.Description.replace(reg, dest));

        var url = document.createElement("p");
        url.className = "url";
        $(url).append(item.DisplayUrl);

        $("#results").append(li);

        var div = document.createElement("div");
        $(li).append(div);
        $(div).append(h3, url, desc);

        $(h3).append(a);
    }

});
