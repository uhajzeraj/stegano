$(document).ready(function() {

    var imgName;
    var imageRow;

    $("button#delete").click(function(event) {

        imageRow = $(event.target).parent().parent().parent()
        imgName = imageRow.find("img").attr("src");

    });
    
    $("button#confirmDelete").click(function() {
        $.ajax({
            url: "deleteImg",
            type: "POST",
            data: {imgName, imgName},
            success: function(data) {
                
                if(data == 1) {
                    imageRow.remove();
                }
    
            }
        });
    });

});