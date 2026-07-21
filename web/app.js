/* 锦书斋网上书店 - 简单前端 SPA */

const API = "/v1";
const PAGE_SIZE = 12;

const state = {
  token: localStorage.getItem("token") || "",
  username: localStorage.getItem("username") || "",
  role: localStorage.getItem("role") || "guest",
  pageNum: 1,
  title: "",
  category: "",
  authMode: "login",
};

/* ---------- 基础工具 ---------- */

function toast(msg, isErr = false) {
  const t = document.getElementById("toast");
  t.textContent = msg;
  t.className = "toast" + (isErr ? " err" : "");
  setTimeout(() => t.classList.add("hidden"), 2600);
}

async function api(path, { method = "GET", body = null, query = null } = {}) {
  let url = API + path;
  if (query) {
    const qs = new URLSearchParams(
      Object.entries(query).filter(([, v]) => v !== "" && v != null)
    ).toString();
    if (qs) url += "?" + qs;
  }
  const headers = { "Content-Type": "application/json" };
  if (state.token) headers["Authorization"] = "Bearer " + state.token;
  const resp = await fetch(url, {
    method,
    headers,
    body: body ? JSON.stringify(body) : null,
  });
  if (resp.status === 403) {
    toast("没有权限，请先登录", true);
    throw new Error("forbidden");
  }
  const data = await resp.json().catch(() => ({ code: -1, message: "响应解析失败" }));
  if (data.code !== 0) {
    toast(data.message || "请求失败", true);
    throw new Error(data.message);
  }
  return data.data;
}

/* ---------- 视图切换 ---------- */

const VIEWS = ["books", "cart", "orders", "manage"];

function showView(name) {
  VIEWS.forEach((v) =>
    document.getElementById("view-" + v).classList.toggle("hidden", v !== name)
  );
  if (name === "books") loadBooks();
  if (name === "cart") loadCart();
  if (name === "orders") loadOrders();
  if (name === "manage") loadManage();
}

/* ---------- 认证 ---------- */

function openAuthModal(mode) {
  state.authMode = mode;
  document.getElementById("auth-title").textContent = mode === "login" ? "登录" : "注册";
  document.getElementById("auth-submit").textContent = mode === "login" ? "登录" : "注册";
  document.getElementById("auth-modal").classList.remove("hidden");
}

function closeAuthModal() {
  document.getElementById("auth-modal").classList.add("hidden");
}

async function submitAuth(e) {
  e.preventDefault();
  const username = document.getElementById("auth-username").value.trim();
  const password = document.getElementById("auth-password").value;
  try {
    if (state.authMode === "register") {
      await api("/user/register", { method: "POST", body: { username, password } });
      toast("注册成功，请登录");
      openAuthModal("login");
      return false;
    }
    const data = await api("/user/login", { method: "POST", body: { username, password } });
    state.token = data.token;
    state.username = username;
    localStorage.setItem("token", state.token);
    localStorage.setItem("username", username);
    closeAuthModal();
    await refreshUserInfo();
    toast("欢迎，" + username);
    showView("books");
  } catch (_) { /* toast 已提示 */ }
  return false;
}

async function refreshUserInfo() {
  try {
    const me = await api("/user/common");
    state.role = me.role || "general";
    localStorage.setItem("role", state.role);
  } catch (_) {
    state.role = "general";
  }
  renderNav();
}

function logout() {
  state.token = "";
  state.username = "";
  state.role = "guest";
  localStorage.removeItem("token");
  localStorage.removeItem("username");
  localStorage.removeItem("role");
  renderNav();
  showView("books");
  toast("已退出登录");
}

function renderNav() {
  const loggedIn = !!state.token;
  const isGeneral = loggedIn && state.role === "general";
  const isManager = loggedIn && (state.role === "seller" || state.role === "admin");
  document.getElementById("btn-login").classList.toggle("hidden", loggedIn);
  document.getElementById("btn-register").classList.toggle("hidden", loggedIn);
  document.getElementById("btn-logout").classList.toggle("hidden", !loggedIn);
  document.getElementById("btn-cart").classList.toggle("hidden", !isGeneral);
  document.getElementById("btn-orders").classList.toggle("hidden", !loggedIn);
  document.getElementById("btn-manage").classList.toggle("hidden", !isManager);
  const info = document.getElementById("user-info");
  info.classList.toggle("hidden", !loggedIn);
  info.textContent = loggedIn ? `👤 ${state.username}（${state.role}）` : "";
}

/* ---------- 书籍浏览 ---------- */

function searchBooks() {
  state.title = document.getElementById("search-input").value.trim();
  state.category = document.getElementById("category-input").value.trim();
  state.pageNum = 1;
  loadBooks();
}

function prevPage() {
  if (state.pageNum > 1) {
    state.pageNum--;
    loadBooks();
  }
}

function nextPage() {
  state.pageNum++;
  loadBooks();
}

async function loadBooks() {
  try {
    const data = await api("/book", {
      query: {
        title: state.title,
        category: state.category,
        pageNum: state.pageNum,
        pageSize: PAGE_SIZE,
      },
    });
    const books = data.bookList || [];
    document.getElementById("page-info").textContent = `第 ${state.pageNum} 页（本页 ${books.length} 本）`;
    renderBooks(books);
    updateCategories(books);
  } catch (_) {}
}

function renderBooks(books) {
  const grid = document.getElementById("book-grid");
  if (!books.length) {
    grid.innerHTML = '<p class="empty">没有找到书籍</p>';
    return;
  }
  grid.innerHTML = books
    .map(
      (b) => `
    <div class="card book-card">
      <h3>${esc(b.title)}</h3>
      <p class="meta">作者：${esc(b.author)}</p>
      <p class="meta">分类：${esc(b.category)} ｜ 出版：${esc(b.publishDate)}</p>
      <p class="price">¥${b.price.toFixed(2)}</p>
      <p class="meta">${b.isSell ? `库存 ${b.number}` : '<span class="off">未上架</span>'}</p>
      ${
        b.isSell && state.role === "general"
          ? `<button class="primary" onclick="addToCart(${b.BookID})">加入购物车</button>`
          : ""
      }
    </div>`
    )
    .join("");
}

function updateCategories(books) {
  const list = document.getElementById("category-list");
  const cats = [...new Set(books.map((b) => b.category).filter(Boolean))];
  const existing = new Set([...list.options].map((o) => o.value));
  cats.forEach((c) => {
    if (!existing.has(c)) {
      const opt = document.createElement("option");
      opt.value = c;
      list.appendChild(opt);
    }
  });
}

/* ---------- 购物车 ---------- */

async function addToCart(bookID) {
  try {
    await api("/cart", { method: "PUT", body: { BookID: bookID, number: 1 } });
    toast("已加入购物车");
  } catch (_) {}
}

async function loadCart() {
  try {
    const data = await api("/cart");
    const books = data.booksInfo || [];
    document.getElementById("cart-price").textContent = (data.cartPrice || 0).toFixed(2);
    const tbody = document.querySelector("#cart-table tbody");
    tbody.innerHTML = books
      .map(
        (b) => `
      <tr>
        <td>${b.index}</td>
        <td>${esc(b.title)}</td>
        <td>¥${b.price.toFixed(2)}</td>
        <td>${b.number}</td>
        <td>¥${(b.price * b.number).toFixed(2)}</td>
        <td><button class="danger" onclick="removeFromCart(${b.BookID}, ${b.number})">移除</button></td>
      </tr>`
      )
      .join("");
  } catch (_) {}
}

async function removeFromCart(bookID, number) {
  try {
    await api("/cart", { method: "DELETE", body: { BookID: bookID, Number: number } });
    toast("已移除");
    loadCart();
  } catch (_) {}
}

async function clearCart() {
  if (!confirm("确定清空购物车？")) return;
  try {
    await api("/cart/all", { method: "DELETE" });
    toast("购物车已清空");
    loadCart();
  } catch (_) {}
}

/* ---------- 订单 ---------- */

async function createOrder() {
  try {
    const data = await api("/order", { method: "POST" });
    toast("下单成功，订单号 " + data.orderId);
    showView("orders");
  } catch (_) {}
}

async function loadOrders() {
  try {
    const data = await api("/order", { query: { pageNum: 1, pageSize: 50 } });
    const orders = data.orderList || [];
    const box = document.getElementById("order-list");
    if (!orders.length) {
      box.innerHTML = '<p class="empty">暂无订单</p>';
      return;
    }
    box.innerHTML = orders
      .map(
        (o) => `
      <div class="card order-card">
        <div class="order-head">
          <strong>订单 #${o.order_id}</strong>
          <span class="status status-${o.status}">${statusText(o.status)}</span>
        </div>
        <p class="meta">下单时间：${esc(o.createdAt)} ｜ 金额：¥${o.orderPrice.toFixed(2)}</p>
        <table class="table">
          <thead><tr><th>书籍ID</th><th>单价</th><th>数量</th></tr></thead>
          <tbody>${(o.orderBook || [])
            .map(
              (b) =>
                `<tr><td>${b.BookID}</td><td>¥${b.unitPrice.toFixed(2)}</td><td>${b.Number}</td></tr>`
            )
            .join("")}</tbody>
        </table>
        ${
          o.status === "open" && state.role === "general"
            ? `<button class="primary" onclick="dealOrder(${o.order_id}, 'accept')">支付</button>
               <button class="danger" onclick="dealOrder(${o.order_id}, 'cancel')">取消</button>`
            : ""
        }
      </div>`
      )
      .join("");
  } catch (_) {}
}

async function dealOrder(orderID, operation) {
  try {
    await api("/order", { method: "PUT", body: { orderId: orderID, operation } });
    toast(operation === "accept" ? "支付成功" : "订单已取消");
    loadOrders();
  } catch (_) {}
}

function statusText(s) {
  return { open: "待支付", accept: "已支付", cancel: "已取消" }[s] || s;
}

/* ---------- 书籍管理（seller/admin） ---------- */

async function addBook(e) {
  e.preventDefault();
  const f = e.target;
  const body = {
    title: f.title.value.trim(),
    author: f.author.value.trim(),
    price: parseFloat(f.price.value),
    publishDate: f.publishDate.value,
    category: f.category.value.trim(),
    isSell: f.isSell.checked,
    number: parseInt(f.number.value, 10),
  };
  try {
    await api("/book", { method: "POST", body });
    toast("添加成功");
    f.reset();
    loadManage();
  } catch (_) {}
  return false;
}

async function loadManage() {
  try {
    const data = await api("/book", { query: { pageNum: 1, pageSize: 100 } });
    const books = data.bookList || [];
    const tbody = document.querySelector("#manage-table tbody");
    tbody.innerHTML = books
      .map(
        (b) => `
      <tr>
        <td>${b.BookID}</td>
        <td>${esc(b.title)}</td>
        <td>${esc(b.author)}</td>
        <td>¥${b.price.toFixed(2)}</td>
        <td>${b.number}</td>
        <td>${b.isSell ? "在售" : "下架"}</td>
        <td>
          <button onclick="toggleSell(${b.BookID}, ${!b.isSell}, '${esc(b.title)}', '${esc(b.author)}', ${b.price}, '${esc(b.publishDate)}', '${esc(b.category)}', ${b.number})">${b.isSell ? "下架" : "上架"}</button>
          <button class="danger" onclick="deleteBook(${b.BookID})">删除</button>
        </td>
      </tr>`
      )
      .join("");
  } catch (_) {}
}

async function toggleSell(id, isSell, title, author, price, publishDate, category, number) {
  try {
    await api("/book/" + id, {
      method: "PUT",
      body: { title, author, price, publishDate, category, isSell, number },
    });
    toast(isSell ? "已上架" : "已下架");
    loadManage();
  } catch (_) {}
}

async function deleteBook(id) {
  if (!confirm("确定删除该书籍？")) return;
  try {
    await api("/book/" + id, { method: "DELETE" });
    toast("已删除");
    loadManage();
  } catch (_) {}
}

/* ---------- 工具 ---------- */

function esc(s) {
  return String(s ?? "").replace(/[&<>"']/g, (c) => ({
    "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#39;",
  }[c]));
}

/* ---------- 启动 ---------- */

(async function init() {
  if (state.token) await refreshUserInfo();
  renderNav();
  showView("books");
})();
