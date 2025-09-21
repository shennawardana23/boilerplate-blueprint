import { expect } from 'chai'
import { setActivePinia, createPinia } from 'pinia'
import { useProjectStore } from '../../src/stores/project.js'

// Mock the API
const mockProjectApi = {
  getTemplates: () => Promise.resolve({ data: { templates: [] } }),
  createProject: () => Promise.resolve({ data: { success: true, project: {} } }),
  getProject: () => Promise.resolve({ data: { success: true, project: {} } }),
  generateProject: () => Promise.resolve({ data: { success: true, files: [] } }),
  downloadProject: () => Promise.resolve({ data: new Blob(), headers: {} })
}

describe('Project Store', () => {
  let store

  beforeEach(() => {
    setActivePinia(createPinia())
    store = useProjectStore()
    
    // Mock the API module
    global.sinon.stub(require('../../src/services/api'), 'projectApi').value(mockProjectApi)
  })

  afterEach(() => {
    global.sinon.restore()
  })

  describe('Initial State', () => {
    it('should have initial state values', () => {
      expect(store.currentProject).to.be.null
      expect(store.projects).to.be.an('array').that.is.empty
      expect(store.templates).to.be.an('array').that.is.empty
      expect(store.isLoading).to.be.false
      expect(store.error).to.be.null
    })

    it('should have computed getters', () => {
      expect(store.hasCurrentProject).to.be.false
      expect(store.currentProjectFiles).to.be.an('array').that.is.empty
      expect(store.isGoProject).to.be.false
      expect(store.isPHPProject).to.be.false
    })
  })

  describe('loadTemplates', () => {
    it('should load templates successfully', async () => {
      const mockTemplates = [
        { language: 'go', name: 'Go Clean Architecture' },
        { language: 'php', name: 'PHP CodeIgniter MVC' }
      ]
      
      mockProjectApi.getTemplates = () => Promise.resolve({ 
        data: { templates: mockTemplates } 
      })

      await store.loadTemplates()

      expect(store.templates).to.deep.equal(mockTemplates)
      expect(store.isLoading).to.be.false
      expect(store.error).to.be.null
    })

    it('should handle load templates error', async () => {
      const errorMessage = 'Failed to load templates'
      mockProjectApi.getTemplates = () => Promise.reject(new Error(errorMessage))

      await store.loadTemplates()

      expect(store.error).to.equal(errorMessage)
      expect(store.isLoading).to.be.false
    })
  })

  describe('createProject', () => {
    it('should create project successfully', async () => {
      const projectData = {
        name: 'test-project',
        language: 'go',
        description: 'A test project'
      }
      
      const mockProject = {
        id: 'test-id',
        ...projectData,
        options: {},
        files: [],
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      }

      mockProjectApi.createProject = () => Promise.resolve({ 
        data: { success: true, project: mockProject } 
      })

      const result = await store.createProject(projectData)

      expect(result).to.deep.equal(mockProject)
      expect(store.currentProject).to.deep.equal(mockProject)
      expect(store.projects).to.include(mockProject)
      expect(store.isLoading).to.be.false
      expect(store.error).to.be.null
    })

    it('should handle create project error', async () => {
      const projectData = { name: 'test-project', language: 'go' }
      const errorMessage = 'Failed to create project'
      
      mockProjectApi.createProject = () => Promise.resolve({ 
        data: { success: false, error: errorMessage } 
      })

      try {
        await store.createProject(projectData)
      } catch (err) {
        expect(err.message).to.equal(errorMessage)
      }

      expect(store.error).to.equal(errorMessage)
      expect(store.isLoading).to.be.false
    })
  })

  describe('loadProject', () => {
    it('should load project successfully', async () => {
      const projectId = 'test-id'
      const mockProject = {
        id: projectId,
        name: 'test-project',
        language: 'go',
        files: []
      }

      mockProjectApi.getProject = () => Promise.resolve({ 
        data: { success: true, project: mockProject } 
      })

      const result = await store.loadProject(projectId)

      expect(result).to.deep.equal(mockProject)
      expect(store.currentProject).to.deep.equal(mockProject)
      expect(store.isLoading).to.be.false
      expect(store.error).to.be.null
    })

    it('should handle load project error', async () => {
      const projectId = 'test-id'
      const errorMessage = 'Project not found'
      
      mockProjectApi.getProject = () => Promise.resolve({ 
        data: { success: false, error: errorMessage } 
      })

      try {
        await store.loadProject(projectId)
      } catch (err) {
        expect(err.message).to.equal(errorMessage)
      }

      expect(store.error).to.equal(errorMessage)
      expect(store.isLoading).to.be.false
    })
  })

  describe('generateProjectFiles', () => {
    it('should generate project files successfully', async () => {
      const projectId = 'test-id'
      const mockFiles = [
        { path: 'go.mod', content: 'module test-project', is_directory: false },
        { path: 'main.go', content: 'package main', is_directory: false }
      ]

      // Set up current project
      store.currentProject = { id: projectId, files: [] }

      mockProjectApi.generateProject = () => Promise.resolve({ 
        data: { success: true, files: mockFiles } 
      })

      const result = await store.generateProjectFiles(projectId)

      expect(result).to.deep.equal(mockFiles)
      expect(store.currentProject.files).to.deep.equal(mockFiles)
      expect(store.isLoading).to.be.false
      expect(store.error).to.be.null
    })

    it('should handle generate project files error', async () => {
      const projectId = 'test-id'
      const errorMessage = 'Failed to generate files'
      
      mockProjectApi.generateProject = () => Promise.resolve({ 
        data: { success: false, error: errorMessage } 
      })

      try {
        await store.generateProjectFiles(projectId)
      } catch (err) {
        expect(err.message).to.equal(errorMessage)
      }

      expect(store.error).to.equal(errorMessage)
      expect(store.isLoading).to.be.false
    })
  })

  describe('downloadProject', () => {
    it('should download project successfully', async () => {
      const projectId = 'test-id'
      const mockBlob = new Blob(['test content'], { type: 'application/zip' })
      const mockHeaders = {
        'content-disposition': 'attachment; filename="test-project-go.zip"'
      }

      mockProjectApi.downloadProject = () => Promise.resolve({ 
        data: mockBlob, 
        headers: mockHeaders 
      })

      // Mock DOM methods
      global.URL.createObjectURL = () => 'blob:test-url'
      global.URL.revokeObjectURL = () => {}
      
      const mockLink = {
        href: '',
        download: '',
        click: () => {},
        remove: () => {}
      }
      global.document.createElement = () => mockLink
      global.document.body.appendChild = () => {}
      global.document.body.removeChild = () => {}

      await store.downloadProject(projectId)

      expect(store.isLoading).to.be.false
      expect(store.error).to.be.null
    })

    it('should handle download project error', async () => {
      const projectId = 'test-id'
      const errorMessage = 'Failed to download project'
      
      mockProjectApi.downloadProject = () => Promise.reject(new Error(errorMessage))

      try {
        await store.downloadProject(projectId)
      } catch (err) {
        expect(err.message).to.equal(errorMessage)
      }

      expect(store.error).to.equal(errorMessage)
      expect(store.isLoading).to.be.false
    })
  })

  describe('updateProjectOptions', () => {
    it('should update current project options', () => {
      const currentOptions = { framework: 'gin', database: 'postgresql' }
      const newOptions = { authentication: 'jwt' }
      const expectedOptions = { ...currentOptions, ...newOptions }

      store.currentProject = {
        id: 'test-id',
        options: currentOptions
      }

      store.updateProjectOptions(newOptions)

      expect(store.currentProject.options).to.deep.equal(expectedOptions)
    })

    it('should not update options if no current project', () => {
      store.currentProject = null

      store.updateProjectOptions({ framework: 'gin' })

      expect(store.currentProject).to.be.null
    })
  })

  describe('clearCurrentProject', () => {
    it('should clear current project', () => {
      store.currentProject = { id: 'test-id', name: 'test-project' }

      store.clearCurrentProject()

      expect(store.currentProject).to.be.null
    })
  })

  describe('clearError', () => {
    it('should clear error', () => {
      store.error = 'Some error message'

      store.clearError()

      expect(store.error).to.be.null
    })
  })

  describe('Computed Properties', () => {
    it('should compute hasCurrentProject correctly', () => {
      expect(store.hasCurrentProject).to.be.false

      store.currentProject = { id: 'test-id' }
      expect(store.hasCurrentProject).to.be.true
    })

    it('should compute currentProjectFiles correctly', () => {
      expect(store.currentProjectFiles).to.be.an('array').that.is.empty

      store.currentProject = { files: [{ path: 'go.mod' }] }
      expect(store.currentProjectFiles).to.deep.equal([{ path: 'go.mod' }])
    })

    it('should compute isGoProject correctly', () => {
      expect(store.isGoProject).to.be.false

      store.currentProject = { language: 'go' }
      expect(store.isGoProject).to.be.true

      store.currentProject = { language: 'php' }
      expect(store.isGoProject).to.be.false
    })

    it('should compute isPHPProject correctly', () => {
      expect(store.isPHPProject).to.be.false

      store.currentProject = { language: 'php' }
      expect(store.isPHPProject).to.be.true

      store.currentProject = { language: 'go' }
      expect(store.isPHPProject).to.be.false
    })
  })
})