const resize = () => {
  const Iwd = window.innerWidth
  let rem = Iwd / 1440
  if (rem <= 0.83) {
    rem = 0.83
  }
  document.documentElement.style.fontSize = rem + 'px'
}

export {
  resize
}
