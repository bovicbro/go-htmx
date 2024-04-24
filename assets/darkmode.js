let mode = localStorage.getItem("mode")

if (mode == "dark") {
  document.getElementById('mainBody')?.classList.toggle('dark')
  document.getElementById('mainBody')?.classList.toggle('light')
}

function darkMode(elementId) {
  document.getElementById(elementId)?.classList.toggle('dark')
  document.getElementById(elementId)?.classList.toggle('light')
  if (mode === "light") {
    localStorage.setItem("mode", "dark")
    mode == "dark"
  }
  if (mode === "dark") {
    localStorage.setItem("mode", "light")
    mode == "light"
  }
  console.log(mode)
}
