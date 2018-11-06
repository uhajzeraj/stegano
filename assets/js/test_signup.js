$(document).ready(function() {
    $(`#register`).click(function(e) {
      
    e.preventDefault();
  
    var formData = new FormData();
      
    // Get the data
    var displayName = $(`#display_name`).val();
    var email = $(`#email`).val();
    var pass = $(`#password`).val();
    var passConfirm = $(`#password_confirmation`).val();
    
    // Append the data
    formData.append('displayName', displayName);
    formData.append('email', email);
    formData.append('pass', pass);
    formData.append('passConfirm', passConfirm);
    
    // Send the data
    $.ajax({
    url: "signup",
    type: "POST",
    data: formData,
    processData: false,
    contentType: false,
    success: function(data) {
        alert(data);
    }
    });
      
    });
  });