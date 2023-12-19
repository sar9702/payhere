// getCookie 함수는 쿠키에서 name에 해당되는 값을 반환하는 함수이다.
function getCookie(name) {
  let match = document.cookie.match(new RegExp("(^| )" + name + "=([^;]+)"));
  if (match) {
    return match[2];
  } else {
    return null;
  }
}

// getURLParam 함수는 URL의 파라미터를 반환하는 함수이다.
function getURLParam(name) {
  let urlSearch = new URLSearchParams(location.search);
  return urlSearch.get(name);
}

// getItemRow 함수는 테이블에 추가될 아이템 행을 반환하는 함수이다.
function getItemRow(item) {
  return `
  <tr>
    <td>${item.Category}</td>
    <td>${item.Name}</td>
    <td>${item.Price}</td>
    <td>${item.Cost}</td>
    <td>
      <a class="badge text-bg-warning" href="/item/detail?id=${item.ID}"
        style="text-decoration-line: none">Detail</a>
      <span class="badge text-bg-danger" onclick="delItem('${item.ID}')"
        style="cursor: pointer">Del</span>
    </td>
  </tr>  
  `;
}

// setInitPage 함수는 메인 페이지가 로드되면 실행되는 함수이다.
function setInitPage() {
  let token = getCookie("SessionToken");

  $.ajax({
    url: `/api/items`,
    type: "get",
    headers: {
      Authorization: "Basic " + token,
    },
    dataType: "json",
    success: function (jsonData) {
      for (let item of jsonData.data) {
        let tr = getItemRow(item);
        $("#itemTable tbody").append(tr);
      }
    },
    error: function (response) {
      alert(
        `code: ${response.responseJSON.meta.code}\nmsg: ${response.responseJSON.meta.message}`
      );
    },
  });
}

// setDetailPage 함수는 아이템 상세 페이지가 로드되면 실행되는 함수이다. 아이템 정보를 가져와 input에 넣어준다.
function setDetailPage() {
  let token = getCookie("SessionToken");
  let id = getURLParam("id");

  $.ajax({
    url: `/api/item/${id}`,
    type: "get",
    headers: {
      Authorization: "Basic " + token,
    },
    dataType: "json",
    success: function (jsonData) {
      let item = jsonData.data;
      $("#category").val(item.Category);
      $("#name").val(item.Name);
      $("#price").val(item.Price);
      $("#cost").val(item.Cost);
      $("#description").text(item.Description);
      $("#barcode").val(item.Barcode);
      $("#expirationDate").val(item.ExpirationDate);
      $("#size option").text(item.Size);
    },
    error: function (response) {
      alert(
        `code: ${response.responseJSON.meta.code}\nmsg: ${response.responseJSON.meta.message}`
      );
    },
  });
}

// delItem 함수는 아이템을 삭제하는 함수이다.
function delItem(id) {
  response = confirm("아이템을 삭제하시겠습니까?");
  if (!response) {
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
