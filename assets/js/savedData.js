$(document).ready(function() {

    $("button#delete").click(function(event) {
        
        imgName = $(event.target).parent().parent().find("img").attr("src");

        $.ajax({
            url: "deleteImg",
            type: "POST",
            data: {imgName, imgName},
            success: function(data) {
              
              if(data == 1) {
                $(event.target).parent().parent().remove();
              }
    
            }
          });
    });
});