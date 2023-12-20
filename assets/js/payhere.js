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

// setInitPage 함수는 메인 페이지가 로드되면 실행되는 함수이다. 테이블에 아이템 행 리스트를 추가한다.
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
function setDetailPage(isEditPage = false) {
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

      if (!isEditPage) {
        // 상세 페이지
        $("#size option").text(item.Size);
        $("#editBtn").attr("href", `/item/edit?id=${item.ID}`);
      } else {
        // 수정 페이지
        $("#size").val(item.Size).attr("selected", true);
      }
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

  let token = getCookie("SessionToken");

  $.ajax({
    url: `/api/item/${id}`,
    type: "delete",
    headers: {
      Authorization: "Basic " + token,
    },
    dataType: "json",
    success: function () {
      alert("아이템이 삭제되었습니다.");
      window.location.reload(); // 페이지 새로고침
    },
    error: function (response) {
      alert(
        `code: ${response.responseJSON.meta.code}\nmsg: ${response.responseJSON.meta.message}`
      );
    },
  });
}

// registerItem 함수는 아이템 등록 페이지에서 등록 버튼을 클릭하면 실행되는 함수이다.
function registerItem() {
  let token = getCookie("SessionToken");

  let sendData = {
    category: $("#category").val(),
    name: $("#name").val(),
    price: $("#price").val(),
    cost: $("#cost").val(),
    description: $("#description").val(),
    barcode: $("#barcode").val(),
    expirationDate: $("#expirationDate").val(),
    size: $("#size option:selected").val(),
  };

  $.ajax({
    url: `/api/item`,
    type: "post",
    headers: {
      Authorization: "Basic " + token,
    },
    data: sendData,
    dataType: "json",
    success: function () {
      // 아이템 등록 완료 페이지로 리다이렉트
      window.location.href = "/item/register-success";
    },
    error: function (response) {
      alert(
        `code: ${response.responseJSON.meta.code}\nmsg: ${response.responseJSON.meta.message}`
      );
    },
  });
}

// editItem 함수는 아이템 수정 페이지에서 수정 버튼을 클릭하면 실행되는 함수이다.
function editItem() {
  let token = getCookie("SessionToken");
  let id = getURLParam("id");

  let sendData = {
    category: $("#category").val(),
    name: $("#name").val(),
    price: $("#price").val(),
    cost: $("#cost").val(),
    description: $("#description").val(),
    barcode: $("#barcode").val(),
    expirationDate: $("#expirationDate").val(),
    size: $("#size option:selected").val(),
  };

  $.ajax({
    url: `/api/item/${id}`,
    type: "put",
    headers: {
      Authorization: "Basic " + token,
    },
    data: sendData,
    dataType: "json",
    success: function () {
      // 아이템 수정 완료 페이지로 리다이렉트
      window.location.href = `/item/edit-success/${id}`;
    },
    error: function (response) {
      alert(
        `code: ${response.responseJSON.meta.code}\nmsg: ${response.responseJSON.meta.message}`
      );
    },
  });
}
