## Asset-Manager high level design

Asset manager is responsible for the managing the DCP/KDM/SPL on the SMS. It is divided in to different components as mentioned below.

### Storage
It store/retrieve the asset manger data on the disk. At high level data is devided in two types.

1. **Assets**  
KDM xml files, SPL xml files and DCPs are the assets. These xml files are maintained in a directory structure as defined below.
   * **kdm** :  Directory to keep kdm xml files on the SMS
   * **spl** : Directory to keep spl xml files on the SMS
   * **dcp** : Directory to keep DCPs on the SMS. Under this directory,  several directories will be created for each DCP, one directory will map to one DCP.   These directories keeps the PKL, CPL, asset-map and mxf files.	            
	      
            

2. **Auxiliary-data**  
      Data which need to  persist across several life cycles of the service. This information would be generated during processing different request. Any DBMS(like sqlite) can be used to maintain this information.
	 

### Asset cache
1. It is introduced in the system to enable faster access to the Assets stored on the disk. This will provide the APIs to store/retrieve the assets and auxiliary-data to/from system.
2. To avoid rebuilding the new cache at every start of service, This cache is persisted on the disk using DBMS(like sqlite).
3. Cache can be rebuilt at any point from the available data in the storage. Cache rebuild might be required in scenarios like software upgrade where older cache format is changed. 
4. If auxiliary-data is corrupted or becomes invalid, then it is ignored during cache rebuild (loss of auxiliary data).
		
### Task manager
1. Task manager gets/stores the assets and auxiliary-data using asset cache.
2. Task manager provides
   1. Interface(APIs) for asset management tasks/operations.
       * Ingest of CPL/SPL/DCP. 
       * Deletion of CPL/SPL/DCP.
       * Get the status of the ongoing asset management tasks/operations.
       * List the Assets and details of particular asset.	
   2. Interface(APIs) for controlling asset management operations (stop the allowing new ingest , etc)   	
   3. Auto delete functionality for CPL/SPL/DCP according to the business rules.
  

### Notification relay
Different components of the asset manager (like task manager, asset cache) will call notification relay API to generate the notifications. These notification will be forwarded the the registered handler. Notifications handler would be

1. Asset Manager Messaging.
2. Logs/Events writer module.

### Messaging:

Asset manager communicates with the external world with messaging layer. It  have follwing components
     
1. **Message/API call/Request-Response** : Handle the messages from external world. When ever it receives any message, it will call  appropriate APIs of the task manager. Depending on the API data, it will construct the response message and send it back to the peer. 
2. **Notification** :
   Mechanism to publish the notifications to external world. Who ever want to get the notification need the subscribe it. It will receive notification from notification relay and forward it to subscriber.
     
