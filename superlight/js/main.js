$(function(){
$("#plus>button").click(function(){
    another=$(".card:first").clone(true)
    another.insertBefore($("#plus"));
    another.find("textarea").val("")
});
$(document).on('click', '.exec>button', function () {
    var card=$(this).parents(".card")
    var text=$(card).find(".code").val()


});
$(document).on('click', '#save', function () {
    var data=[]
    $(".card textarea.code").each(function(){
        data.push($(this).val())
    })
    console.log(data)

    $.post("/save",{"data":data})

});

})