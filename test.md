## Asset-Manager high level design

Asset manager is responsible for the managing the DCP/KDM/SPL on the SMS. It is divided in to different components as mentioned below.

### Storage
It store the asset manger data on the disk. At high level data is divided in two types.

1. **Assets**  
KDM, SPL and DCPs are the assets. This is maintained in a directory structure as defined below.
   * **kdm** :  Directory to keep kdm on the SMS
   * **spl** : Directory to keep spl on the SMS
   * **dcp** : Directory to keep DCPs on the SMS. Under this directory,  several directories will be created for each DCP, one directory will map to one DCP.   These directories keeps the PKL, CPL, asset-map and media files.	            	    
2. **DB**

   DB (sqlite) is mainly divided in to parts

   * *Auxiliary-data*  
      Data which need to  persist across several life cycles of the service. This information would be generated during processing different request.
      
      e.g. 
      * Ingest status information
      * Protect the CPL from GC
      * ....
      
   * *Asset cache*  
     It will keep the part of the information in the assets files on the disk . It is introduced in the system to enable faster access to the assets information.
     
     e.g.
     * KDM information
     * CPL information 
     * ....
     
     
   DB can be rebuilt at any point from the available assets in the storage. DB rebuild/upgrade might be required in scenarios like software upgrade where older DB format is changed. If auxiliary-data becomes invalid or unusable, then it is ignored during DB rebuild (loss of auxiliary-data in DB).

### Asset store
This will provide the APIs to other components in asset manager to store/retrieve/access the assets and DB. It provides storage abstraction layer.

### Task manager
The Task manager manages the scheduling of tasks based on their categories, types and priorities. It will decide which tasks can run in parallel or cannot run due to external triggers(restrictions). It will persist its task queue to the data store.

e.g. A playback trigger may prevent ingest tasks from running.

#### Task
Tasks are defined as the entities that will implement the functionality defined by the API.
Tasks would include:

  * SPL/KDM/DCP Ingest
  * Fetch Ingest queue
  * Suspend/Resume/Delete ingest
  * Fetch SPL/KDM/DCP List
  * Fetch SPL/KDM/DCP information
  * Delete SPL/KDM/DCP
  * Asset auto deletion 
  * ...etc  

### Notification relay
Different components of the asset manager (like task manager, asset cache) will call notification relay API to generate the notifications. These notification will be forwarded the the registered handler. Notifications handler would be

1. Asset Manager Messaging.
2. Logs/Events writer module.

### Messaging:

Asset manager communicates with the external world with messaging layer. It  have follwing components
     
1. **Message/API call/Request-Response** : Handle the messages from external world. When ever it receives any message, it will call  appropriate APIs of the task manager. Depending on the API data, it will construct the response message and send it back to the peer. 
2. **Notification** :
   Mechanism to publish the notifications to external world. Who ever want to get the notification need the subscribe it. It will receive notification from notification relay and forward it to subscriber.
     
