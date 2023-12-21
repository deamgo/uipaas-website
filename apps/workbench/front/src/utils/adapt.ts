const resize = () => {
  const Iwd = window.innerWidth
  let rem = Iwd / 1440
  if (rem <= 0.8) {
    rem = 0.8
  }
  document.documentElement.style.fontSize = rem + 'px'
}

export {
  resize
}
