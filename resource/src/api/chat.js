import request from '@/utils/request'

// 提交表单信息
export function submitQuestion(data) {
    return request({
      url: '/api/chat',
      method: 'post',
      data
    })
  }