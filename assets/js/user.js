
$(document).ready(function() {
    
    $("#changePass").click(function() {
        var oldPass = $("#currentPass").val();
        var newPass = $("#newPass").val();
        var confirmPass = $("#confirmPass").val();

        var formData = new FormData();
        // Append the data
        formData.append('currentPass', oldPass);
        formData.append('newPass', newPass);
        formData.append('confirmPass', confirmPass);

        // Send the data
        $.ajax({
            url: "changePass",
            type: "POST",
            data: formData,
            processData: false,
            contentType: false,
            success: function(data) {
                
                if(data == 1) {
                    $("#danger-row").fadeOut("fast");
                    $(".alert-success").fadeIn("fast");
                } else { // We got some errors

                    $("#danger-row").fadeIn("fast");
                    var output = "";
                    var jsonData = JSON.parse(data);

                    for(var i = 0; i < jsonData.length; i++) {
                      output += jsonData[i] + "<br/>";
                    }
                    $("#danger-info").html(output);
                }

            }
        });

    });  

    // Delete account
    $("#deleteAcc").click(function() {
        $.ajax({
            url: "deleteAcc",
            type: "DELETE",
            success: function(data) {
                if(data == 1) {
                    window.location.replace("/");
                }
            }
        });
    });

    // Danger alert close button won't work, close it by this
    $('#danger-row button.close').click(function() {
        $("#danger-row").fadeOut("fast");
    });

    // Danger alert close button won't work, close it by this
    $('.alert-success button.close').click(function() {
        $(".alert-success").fadeOut("fast");
    });

});