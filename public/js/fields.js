$(document).ready(function(){
    var next = 1;
    $(".add-more").click(function(e){
        e.preventDefault();
        var addto = "#loc" + next;
        var addRemove = "#loc" + (next);
        next = next + 1;
        var newIn = '<select id="loc' + next + '" name="loc' + next + '" class="form-control">';
        newIn +=
        '<option selected>Choose...</option>\
            <option>Ahmedguda</option>\
            <option>Allwyn Colony</option>\
            <option>Ameenpur</option>\
            <option>Ameerpet</option>\
            <option>Bachupally</option>\
            <option>Begumpet</option>\
            <option>Borabanda</option>\
            <option>Bowenpally</option>\
            <option>Bhel</option>\
            <option>Brundavan Colony</option>\
            <option>Chandanagar</option>\
            <option>Fatehnagar</option>\
            <option>Gachibowli</option>\
            <option>Gautam Nagar</option>\
            <option>Hafeezpet</option>\
            <option>Hitec City</option>\
            <option>Hydernagar</option>\
            <option>Jubilee Hills</option>\
            <option>Kandi</option>\
            <option>Kompally</option>\
            <option>Kondapur</option>\
            <option>Kukatpally</option>\
            <option>Lingampally</option>\
            <option>Madhapur</option>\
            <option>Miyapur</option>\
            <option>Nallagandla</option>\
            <option>Nizampet</option>\
            <option>Patancheru</option>\
            <option>Patel Nagar</option>\
            <option>Sai Nagar</option>\
            <option>Sanath Nagar</option>\
            <option>Sangareddy</option>\
            <option>Shamshiguda</option>\
            <option>Siva Nagar</option>\
            <option>Yeddumailaram</option>\
          </select>';
        var newInput = $(newIn);
        var removeBtn = '<button id="remove' + (next - 1) + '" class="btn btn-danger remove-me" >-</button>';
        var removeButton = $(removeBtn);
        $(addto).after(newInput);
        $(addRemove).after(removeButton);
        $("#loc" + next).attr('data-source',$(addto).attr('data-source'));
        $('.remove-me').click(function(e){
            e.preventDefault();
            var fieldNum = this.id.charAt(this.id.length-1);
            var fieldID = "#loc" + fieldNum;
            $(this).remove();
            $(fieldID).remove();
        });
    });
});
