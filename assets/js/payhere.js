// delItem 함수는 아이템을 삭제하는 함수이다.
function delItem(id) {
  response = confirm("아이템을 삭제하시겠습니까?");
  if (!response) {
    console.log("아니오");
    return;
  }

  $.ajax({
    url: `/item/${id}`,
    type: "delete",
    success: function () {
      alert("아이템이 삭제되었습니다.");
      window.location.reload(); // 페이지 새로고침
    },
    error: function (request, status, error) {
      alert(
        `code: ${request.status}\nstatus: ${status}\nmsg: ${request.responseText}\nerror: ${error}`
      );
    },
  });

  console.log(id);
}
