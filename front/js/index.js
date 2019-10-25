let api = "http://127.0.0.1:12345/api"

$(document).ready(function() {
    $("#url-form").submit(function(e){
        e.preventDefault();
        $.ajax({
            url : api + "?url=" + $("#url").val(),
            type : 'GET',
            dataType:'json',
            success : function(data) {    
                $("#notif").fadeOut();          
                $("#arrow").fadeOut().fadeIn();
                $("#result").fadeOut().delay(1000).fadeIn();
                $("#result").text(data.short);

            },
            error : function(jqXHR, textStatus, errorThrown)
            {
                $("#notif").fadeOut().fadeIn();
                $("#notif").text(textStatus);
            }
        });
    });
});