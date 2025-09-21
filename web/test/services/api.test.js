import { expect } from 'chai'
import axios from 'axios'
import { projectApi, chatApi, healthApi } from '../../src/services/api.js'

// Mock axios
const mockAxios = {
  get: () => Promise.resolve({ data: {} }),
  post: () => Promise.resolve({ data: {} }),
  interceptors: {
    request: { use: () => {} },
    response: { use: () => {} }
  }
}

describe('API Services', () => {
  describe('projectApi', () => {
    it('should have getTemplates method', () => {
      expect(projectApi.getTemplates).to.be.a('function')
    })

    it('should have createProject method', () => {
      expect(projectApi.createProject).to.be.a('function')
    })

    it('should have getProject method', () => {
      expect(projectApi.getProject).to.be.a('function')
    })

    it('should have generateProject method', () => {
      expect(projectApi.generateProject).to.be.a('function')
    })

    it('should have downloadProject method', () => {
      expect(projectApi.downloadProject).to.be.a('function')
    })

    describe('getTemplates', () => {
      it('should make GET request to /templates', async () => {
        const mockResponse = { data: { templates: [] } }
        const axiosSpy = global.sinon.stub(axios, 'get').resolves(mockResponse)
        
        await projectApi.getTemplates()
        
        expect(axiosSpy.calledWith('/api/templates')).to.be.true
        axiosSpy.restore()
      })
    })

    describe('createProject', () => {
      it('should make POST request to /projects with project data', async () => {
        const projectData = { name: 'test-project', language: 'go' }
        const mockResponse = { data: { success: true, project: projectData } }
        const axiosSpy = global.sinon.stub(axios, 'post').resolves(mockResponse)
        
        await projectApi.createProject(projectData)
        
        expect(axiosSpy.calledWith('/api/projects', projectData)).to.be.true
        axiosSpy.restore()
      })
    })

    describe('getProject', () => {
      it('should make GET request to /projects/:id', async () => {
        const projectId = 'test-id'
        const mockResponse = { data: { success: true, project: {} } }
        const axiosSpy = global.sinon.stub(axios, 'get').resolves(mockResponse)
        
        await projectApi.getProject(projectId)
        
        expect(axiosSpy.calledWith(`/api/projects/${projectId}`)).to.be.true
        axiosSpy.restore()
      })
    })

    describe('generateProject', () => {
      it('should make POST request to /projects/:id/generate', async () => {
        const projectId = 'test-id'
        const mockResponse = { data: { success: true, files: [] } }
        const axiosSpy = global.sinon.stub(axios, 'post').resolves(mockResponse)
        
        await projectApi.generateProject(projectId)
        
        expect(axiosSpy.calledWith(`/api/projects/${projectId}/generate`)).to.be.true
        axiosSpy.restore()
      })
    })

    describe('downloadProject', () => {
      it('should make GET request to /projects/:id/download with blob response type', async () => {
        const projectId = 'test-id'
        const mockResponse = { data: new Blob(), headers: {} }
        const axiosSpy = global.sinon.stub(axios, 'get').resolves(mockResponse)
        
        await projectApi.downloadProject(projectId)
        
        expect(axiosSpy.calledWith(`/api/projects/${projectId}/download`, {
          responseType: 'blob'
        })).to.be.true
        axiosSpy.restore()
      })
    })
  })

  describe('chatApi', () => {
    it('should have sendMessage method', () => {
      expect(chatApi.sendMessage).to.be.a('function')
    })

    it('should have getChatHistory method', () => {
      expect(chatApi.getChatHistory).to.be.a('function')
    })

    describe('sendMessage', () => {
      it('should make POST request to /chat/message with message data', async () => {
        const messageData = { message: 'Hello', project_id: 'test-id' }
        const mockResponse = { data: { success: true, message: {} } }
        const axiosSpy = global.sinon.stub(axios, 'post').resolves(mockResponse)
        
        await chatApi.sendMessage(messageData)
        
        expect(axiosSpy.calledWith('/api/chat/message', messageData)).to.be.true
        axiosSpy.restore()
      })
    })

    describe('getChatHistory', () => {
      it('should make GET request to /chat/history without params when no projectId', async () => {
        const mockResponse = { data: { success: true, history: {} } }
        const axiosSpy = global.sinon.stub(axios, 'get').resolves(mockResponse)
        
        await chatApi.getChatHistory()
        
        expect(axiosSpy.calledWith('/api/chat/history', { params: {} })).to.be.true
        axiosSpy.restore()
      })

      it('should make GET request to /chat/history with project_id param', async () => {
        const projectId = 'test-id'
        const mockResponse = { data: { success: true, history: {} } }
        const axiosSpy = global.sinon.stub(axios, 'get').resolves(mockResponse)
        
        await chatApi.getChatHistory(projectId)
        
        expect(axiosSpy.calledWith('/api/chat/history', { params: { project_id: projectId } })).to.be.true
        axiosSpy.restore()
      })
    })
  })

  describe('healthApi', () => {
    it('should have check method', () => {
      expect(healthApi.check).to.be.a('function')
    })

    describe('check', () => {
      it('should make GET request to /health', async () => {
        const mockResponse = { data: { status: 'healthy' } }
        const axiosSpy = global.sinon.stub(axios, 'get').resolves(mockResponse)
        
        await healthApi.check()
        
        expect(axiosSpy.calledWith('/api/health')).to.be.true
        axiosSpy.restore()
      })
    })
  })
})