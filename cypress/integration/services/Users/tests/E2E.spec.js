import * as PUTUsers from '../requests/PUTUsers.request'
import * as GETUsers from '../requests/GETUsers.request'
import * as GETUser from '../requests/GETUser.request'
import * as POSTUsers from '../requests/POSTUsers.request'
import * as DELETEUsers from '../requests/DELETEUser.request'

describe('E2E Users', () => {
    var responseUser = null
it('Add a new user and check the users list', () => {
    POSTUsers.addUser().then((response) => {
        GETUsers.allUsers().should((respGet) => {
            respGet.body.data.forEach(u => {
              if (u.name === response.name) response.gender = u.gender
            })
            expect(respGet.gender).to.be.equal('Male')
            expect(respGet.status).to.eq(200);
            expect(respGet.body).to.be.not.null;
        })
      expect(response.status).to.eq(200);
      expect(response.body.data.name).to.equal("User Test");
      expect(response.body.data.email).to.be.not.null;
      expect(response.body.data.gender).to.equal("Male");

      responseUser = response

    })
  });

  it('Update a user and get user updated by id', () => {
    PUTUsers.updateUser(responseUser.body.data.id).then((response) => {
        GETUser.getUser(responseUser.body.data.id).should((respGet) => {
            expect(respGet.body.data.name).to.be.equal('User Test Update')
            expect(respGet.status).to.eq(200);
        })
      expect(response.status).to.eq(200);
      expect(response.body.data.name).to.equal("User Test Update");
    })
  });

  it('Delete a user and get user updated by id', () => {
    DELETEUsers.deleteUSer(responseUser.body.data.id).then((response) => {
        GETUser.getUser(responseUser.body.data.id).should((respGet) => {
            expect(respGet.status).to.eq(404);
        })
      expect(response.status).to.eq(200);
    })
  });

})