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
                }
            }
        });
          
    });

});