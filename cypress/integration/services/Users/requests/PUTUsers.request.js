/// <reference types="cypress" />
const payloadUpdateUser = require('../payloads/update-user.json')
const token = Cypress.env('token')

function updateUser(idUser) {
  return cy.request({
    method: 'PUT',
    url: `/users/${idUser}`,
    body: payloadUpdateUser,
    auth:{
            bearer: token
        } 
  })
}

export { updateUser };