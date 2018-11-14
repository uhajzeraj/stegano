$(document).ready(function() {

    $(`button#loginbtn`).click(function(e) {
      
        e.preventDefault();

        var formData = new FormData();
          
        // Get the data
        var user = $(`#user`).val();
        var pass = $(`#pwd`).val();
        
        // Append the data
        formData.append('username', user);
        formData.append('pass', pass);
        
        // Send the data
        $.ajax({
            url: "login",
            type: "POST",
            data: formData,
            processData: false,
            contentType: false,
            success: function(data) {
                
                if(data == 1) {
                    window.location.replace("home");
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
    
    // Danger alert close button won't work, close it by this
    $('button.close').click(function() {
        $("#danger-row").fadeOut("fast");
    });

});