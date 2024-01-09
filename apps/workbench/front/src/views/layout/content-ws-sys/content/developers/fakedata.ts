const makeFake = () => {
  return new Array(88).fill(0).map((index) => {
    return {
      developer_id: 'plokos' + index,
      username: String(index),
      email: index + '@fake.cop',
      role: index,
      status: index,
    }
  })
}

export const fakeData = makeFake()