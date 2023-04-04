// On load
$(document).ready(function() {
    // get the deadline inpt and set the value to the current date
    $('#deadline').val(new Date().toISOString().substr(0, 10));

        // oncilck any  of plus-buttons

});


function onPlus(day) {
    // add day value to the deadline input
    var deadline = $('#deadline').val();
    var date = new Date(deadline);
    date.setDate(date.getDate() + day);
    $('#deadline').val(date.toISOString().substr(0, 10));

    return false
}